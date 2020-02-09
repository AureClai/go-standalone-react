import React, { useEffect, useState, Component } from "react";
import logo from "./logo.svg";
import "./App.css";

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      isWSConnected: false
    };
  }

  componentDidMount() {
    let ws = new WebSocket("ws://localhost:8080/ws");
    let isWSConnected = true;
    this.setState({ isWSConnected });
  }
  // Function for messaging
  theMessage() {
    let { isWSConnected } = this.state;
  }

  render() {
    let { isWSConnected } = this.state;
    let theMessage = "Pas de connection Websocket";
    if (isWSConnected) {
      theMessage = "La connection Websocket est bonne";
    }

    return (
      <div className="App">
        <header className="App-header">
          <img src={logo} className="App-logo" alt="logo" />
          <p>{theMessage}</p>
          <a
            className="App-link"
            href="https://reactjs.org"
            target="_blank"
            rel="noopener noreferrer"
          >
            Learn React
          </a>
        </header>
      </div>
    );
  }
}

export default App;
