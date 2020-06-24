import React from "react";
import ReactDOM from "react-dom";
import "core-js/stable";

import "./css/lato.css";
import "./css/wally.css";
import App from "./components/App";

import Wails from "@wailsapp/runtime";


class Wally extends React.Component {
  state = {
    ready: false,
    step: 0
  }

  componentDidMount() {
    window.wails.Events.On("state_update", state => {
      this.setState({...state, ready: true})
    })
  }

  render() {
    return <App state={this.state} />;
  }
}

Wails.Init(() => {
  ReactDOM.render(<Wally />, document.getElementById("app"));
});
