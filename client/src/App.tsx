import React, { useEffect } from 'react';
import logo from './logo.svg';
import './App.css';
import { Header, Container } from 'semantic-ui-react';
import MainForm from "./mainForm";

const style = {
  h1: {
    marginTop: '3em',
  },
}

const App = () => {
  return (
    <div className="App">
    <Header as='h1' content='Entangle' style={style.h1} textAlign='center' />
    <img src={logo} className="App-logo" alt="logo" />
    <Container>
      <MainForm />
    </Container>
  </div>
  );
}

export default App;
