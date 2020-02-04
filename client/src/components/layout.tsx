import React, { useEffect } from 'react';
import { Header, Container } from 'semantic-ui-react';
import logo from '../logo.svg';

const style = {
    h1: {
        marginTop: '3em',
    },
}

const App = ({children}) => {
    return (
        <>
            <Header as='h1' content='Entangle' style={style.h1} textAlign='center' />
            <img src={logo} className="App-logo" alt="logo" />
            <Container>
                {children}
            </Container>
        </>
    );
}

export default App;