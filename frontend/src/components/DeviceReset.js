import React from "react";
import PlanckReset from "../images/planck-reset.png";
import ErgodoxReset from "../images/ergodox-reset.png";

export default props => {
  const { model } = props.device;
  return (
    <div>
      <div className="media-container">
        <div className="media offset">
          <img
            alt="Flash"
            className="glyph"
            src={model === 0 ? PlanckReset : ErgodoxReset}
          />
        </div>
      </div>
      <h3>Press your keyboard's reset button</h3>
      {model === 0 && (
        <p>
          You’re going to need a paperclip for this: The reset button is located
          at the top left of the back of your keyboard.
        </p>
      )}
      {model === 1 && (
        <p>
          You’re going to need a paperclip for this: The reset button is located
          on the right half of your keyboard, next to the three LEDs.
        </p>
      )}
      {model === 2 && (
        <p>
          You’re going to need a paperclip for this: The reset button is located
          on the left half of your keyboard, next to the three LEDs.
        </p>
      )}
    </div>
  );
};
