package services

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/google/uuid"
	"github.com/primekobie/lucy/internal/mailer"
	"github.com/primekobie/lucy/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	store models.UserStore
	mail  *mailer.Mailer
}

func NewUserService(store models.UserStore, mailer *mailer.Mailer) *UserService {
	return &UserService{
		store: store,
		mail:  mailer,
	}
}

func (us *UserService) CreateUser(ctx context.Context, name, email, password string) (*models.User, error) {

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		slog.Error("failed to hash password", "error", err)
		return nil, ErrFailedOperation
	}
	now := time.Now().UTC()
	user := &models.User{
		ID:           uuid.New(),
		Name:         name,
		Email:        email,
		PasswordHash: hash,
		CreatedAt:    now,
		LastModified: now,
		Verified:     false,
	}

	if err := us.store.Create(ctx, user); err != nil {
		return nil, err
	}

	otpString := generateOTP()
	slog.Debug("OTP verificatio code", "code", otpString) //TODO: delete this line later
	otpHash := hashString(otpString)

	userAddr := mailer.Address{Name: user.Name, Email: user.Email}
	data := mailer.Data{
		Address: userAddr,
		Code:    otpString,
	}

	token := models.UserToken{
		Hash:      otpHash,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Scope:     VERIFICATION,
	}

	_ = us.store.InsertToken(ctx, &token)

	us.sendEmail([]mailer.Address{userAddr}, "verify_email.gohtml", data)

	return user, nil
}

func (us *UserService) VerifyUser(ctx context.Context, code string, email string) (models.User, error) {

	hash := hashString(code)
	user, err := us.store.GetUserForToken(ctx, hash, VERIFICATION, email)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return models.User{}, ErrInvalidToken
		}
		return models.User{}, err
	}

	user.Verified = true

	err = us.store.Update(ctx, &user)
	if err != nil {
		return models.User{}, err
	}

	// Delete otp after successful verification
	_ = us.store.DeleteToken(ctx, hash, VERIFICATION)

	address := mailer.Address{Name: user.Name, Email: user.Email}
	us.sendEmail([]mailer.Address{address}, "welcome_email.gohtml", mailer.Data{Address: address})

	return user, nil
}

func (us *UserService) ResendVerificationEmail(ctx context.Context, email string) error {
	user, err := us.store.GetByEmail(ctx, email)
	if err != nil {
		return err
	} else if user.Verified {
		return errors.New("user already verified")
	}

	otpString := generateOTP()
	otpHash := hashString(otpString)

	slog.Debug("OTP verification code", "code", otpString) //TODO: delete this line later

	userAddr := mailer.Address{Email: email, Name: user.Name}
	data := mailer.Data{
		Address: userAddr,
		Code:    otpString,
	}

	token := models.UserToken{
		Hash:      otpHash,
		UserId:    user.ID,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Scope:     VERIFICATION,
	}

	err = us.store.InsertToken(ctx, &token)
	if err != nil {
		return err
	}

	us.sendEmail([]mailer.Address{userAddr}, "verify_email.gohtml", data)

	return nil
}

func (us *UserService) NewSession(ctx context.Context, email string, password string) (*UserSession, error) {
	user, err := us.store.GetByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	if !user.Verified {
		return nil, ErrUnverifiedUser
	}

	err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return nil, ErrInvalidCredentials
		}
		slog.Error("failed to compare password and hash", "error", err.Error())
		return nil, ErrFailedOperation
	}

	refreshttl := 15 * (24 * time.Hour)
	refresh, err := GenerateToken(user.ID, user.Email, refreshttl, TokenTypeRefresh)
	if err != nil {
		return nil, ErrFailedOperation
	}

	accessttl := 1 * time.Hour
	access, err := GenerateToken(user.ID, user.Email, accessttl, TokenTypeRefresh)
	if err != nil {
		return nil, ErrFailedOperation
	}

	refreshExpiry := time.Now().Add(refreshttl)
	tokenHash := hashString(refresh)
	token := models.UserToken{
		Hash:      tokenHash,
		UserId:    user.ID,
		ExpiresAt: refreshExpiry,
		Scope:     AUTHENTICATION,
	}

	err = us.store.InsertToken(ctx, &token)
	if err != nil {
		return nil, err
	}

	session := &UserSession{
		User:             *user,
		RefreshToken:     refresh,
		RefreshExpiresAt: refreshExpiry,
		AccessToken:      access,
		AccessExpiresAt:  time.Now().Add(accessttl),
	}

	return session, nil
}

func (us *UserService) RefreshSession(ctx context.Context, refreshToken string) (*UserAccess, error) {
	claims, err := ValidateToken(refreshToken, TokenTypeRefresh)
	if err != nil {
		slog.Error("failed token validation", "error", err.Error())
		return nil, ErrInvalidToken
	}

	hash := hashString(refreshToken)

	user, err := us.store.GetUserForToken(ctx, hash, AUTHENTICATION, claims.Email)
	if err != nil {
		if errors.Is(err, models.ErrNotFound) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	ttl := 2 * time.Hour // TODO: make time shorter
	accessToken, err := GenerateToken(user.ID, user.Email, ttl, TokenTypeAccess)
	if err != nil {
		return nil, err
	}
	// FIXME: obtain expiry time from GenerateToken function
	useracc := &UserAccess{
		AccessToken: accessToken,
		ExpiresAt:   time.Now().Add(ttl),
	}

	return useracc, nil
}

func (us *UserService) UpdateUser(ctx context.Context, userData map[string]any) (*models.User, error) {
	id, ok := userData["id"]
	if !ok {
		return nil, errors.New("user id not found")
	}

	user, err := us.store.GetByID(ctx, id.(uuid.UUID))
	if err != nil {
		return nil, err
	}

	name, ok := userData["name"]
	if ok {
		user.Name = name.(string)
	}

	profilePhoto, ok := userData["profilePhoto"]
	if ok {
		user.ProfilePhoto = profilePhoto.(string)
	}

	password, ok := userData["password"]
	if ok {
		if len(password.(string)) < 8 || len(password.(string)) > 20 {
			return nil, ErrInvalidPassword
		}
		err = bcrypt.CompareHashAndPassword(user.PasswordHash, []byte(password.(string)))
		if err != nil {
			if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
				hash, err := bcrypt.GenerateFromPassword([]byte(password.(string)), bcrypt.DefaultCost)
				if err != nil {
					return nil, ErrFailedOperation
				}

				user.PasswordHash = hash
			} else {
				slog.Error("failed to compare password and hash", "error", err.Error())
				return nil, ErrFailedOperation
			}
		}

	}

	user.LastModified = time.Now().UTC()

	err = us.store.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) FetchUser(ctx context.Context, id uuid.UUID) (*models.User, error) {
	user, err := s.store.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	return s.store.Delete(ctx, id)
}
