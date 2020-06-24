import React from "react";

import ZSALogo from "../images/zsa-logo.png";
import ErgodoxLogo from "../images/ergodox-logo.svg";
import PlanckLogo from "../images/planck-logo.svg";

export default ({ appVersion, model, step, toggleLog, hasError }) => {
  const hasLogo = model < 2;
  return (
    <div className="footer">
      {hasLogo && <div className="title">KEYBOARD:</div>}
      <div className="status">
        {step === 0 && "LOOKING..."}
        {step === 1 && "-SELECT-"}
        {step > 1 && hasLogo && (
          <img
            alt="Planck Logo"
            className="logo"
            src={model === 0 ? PlanckLogo : ErgodoxLogo}
          />
        )}
        {(step === 2 || step === 3) && (
          <a
            className="reset"
            href="#reset"
            onClick={e => {
              e.preventDefault();
              window.backend.State.ResetState();
            }}
          >
            Restart
          </a>
        )}
      </div>
      <div className="log-toggle" onClick={toggleLog}>
        >_
        {hasError === true && <span className="bubble">!</span>}
      </div>
      <div>
        <img alt="ZSA" className="logo zsa-logo" src={ZSALogo} />
      </div>
      <div className="version">V{appVersion}</div>
    </div>
  );
};
