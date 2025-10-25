/* @refresh reload */
import { render } from "solid-js/web";
import "./index.css";
import { Router } from "@solidjs/router";
import routes from "./router";

const wrapper = document.getElementById("root");

if (!wrapper) {
  throw new Error("Wrapper div not found");
}

render(() => <Router>{routes}</Router>, wrapper);
