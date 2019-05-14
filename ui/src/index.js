import { preact, h, render } from "preact";
import "./css/lato.css";
import "./css/wally.css";
import App from "./components/App";

const elRoot = document.createElement("div");
document.body.appendChild(elRoot);

const renderLoop = () =>
  render(<App state={state} />, elRoot, elRoot.lastElementChild);

state.render = renderLoop;

renderLoop();
