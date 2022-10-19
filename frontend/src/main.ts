import "./css/inter.css";
import "./css/global.css";
import App from "./App.svelte";

import "./lib/state.ts";

const app = new App({
  target: document.getElementById("app"),
});

export default app;
