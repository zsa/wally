import React from "react";
import SearchGlyph from "../images/search.svg";

export default class DeviceSelect extends React.Component {
  state = {
    ticks: 0
  };

  componentDidMount() {
    this.startProbing();
  }

  componentWillUnmount() {
    this.stopProbing();
  }

  startProbing = () => {
    this.probeInterval = setInterval(() => {
      const { ticks } = this.state;
      this.setState({ ticks: ticks + 1 });
    }, 1000);
  };

  stopProbing = e => {
    clearInterval(this.probeInterval);
  };

  render() {
    const { ticks } = this.state;
      return (
        <div>
          <div className="media-container">
            <div className="media">
              <img alt="Search" className="glyph" src={SearchGlyph} />
            </div>
          </div>
          <h3>Looking for compatible keyboards</h3>
          {ticks < 3 && <p>Hold on while your keyboard is being detected</p>}
          {ticks >= 3 && <p>Make sure your keyboard is firmly connected</p>}
        </div>
      );
  }
}
