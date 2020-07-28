import React from "react";

import ZSALogo from "../images/zsa-logo.png";

export default ({ appVersion, model, step, toggleLog, hasError }) => {
  let modelLabel;
  switch (model) {
    case 0:
      modelLabel = "PLANCK EZ";
      break;
    case 1:
      modelLabel = "ERGODOX EZ";
      break;
    case 2:
      modelLabel = "MOONLANDER MK1";
      break;
    default:
      modelLabel = "";
      break;
  }
  return (
    <div className="footer">
      <div className="title">KEYBOARD:</div>
      <div className="status">
        <div className="model">
          {step === 0 && "LOOKING..."}
          {step === 1 && "-SELECT-"}
          {step > 1 && modelLabel}
        </div>
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
