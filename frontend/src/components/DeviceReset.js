import React from "react";
import PlanckReset from "../images/planck-reset.png";
import ErgodoxReset from "../images/ergodox-reset.png";
import MoonlanderReset from "../images/moonlander-reset.png";

export default props => {
  const { model } = props.device;
  let glyph;
  let glyphOffset = true;
  switch (model) {
    case 0:
      glyph = PlanckReset;
      break;
    case 1:
      glyph = ErgodoxReset;
      break;
    case 2:
      glyph = MoonlanderReset;
      glyphOffset = false;
      break;
    default:
      glyph = null;
      break;
  }

  return (
    <div>
      <div className="media-container">
        <div className={`media${glyphOffset === true && " offset"}`}>
          {glyph && (
            <img alt="Reset your board" className="glyph" src={glyph} />
          )}
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
