import { Component, h } from "preact";

import ZSALogo from "../images/zsa-logo.png";
import ErgodoxLogo from "../images/ergodox-logo.svg";
import PlanckLogo from "../images/planck-logo.svg";

export default ({ model, step, toggleLog, hasError }) => {
  return (
    <div className="footer">
      <div className="title">KEYBOARD:</div>
      <div className="status">
        {step === 0 && "LOOKING..."}
        {step === 1 && "-SELECT-"}
        {step > 1 && (
          <img
            alt="Planck Logo"
            className="logo"
            src={model === 0 ? PlanckLogo : ErgodoxLogo}
          />
        )}
      </div>
      <div className="log-toggle" onClick={toggleLog}>
        >_
        {hasError == true && <span className="bubble">!</span>}
      </div>
      <div>
        <img alt="ZSA" className="logo zsa-logo" src={ZSALogo} />
      </div>
      <div className="version">V1.0.0</div>
    </div>
  );
};
