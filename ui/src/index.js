import React from "react";
import { render } from "react-dom";
import "core-js/stable";

import "./css/lato.css";
import "./css/wally.css";
import App from "./components/App";

const elRoot = document.createElement("div");
document.body.appendChild(elRoot);

const renderLoop = () => render(<App state={state} />, elRoot);

state.render = renderLoop;

renderLoop();
