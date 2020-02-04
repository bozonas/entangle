import React from 'react';
import { BrowserRouter, Switch, Route } from 'react-router-dom'
import './App.css';
import Layout from "./components/layout";
import MainForm from "./components/mainForm";
import DisplaySecret from "./components/displaySecret";
import NotFound from "./components/notFound";
import ErrorPage from "./components/errorPage";

const App = () => {
  return (
    <BrowserRouter>
      <div className="App">
        <Layout>
          <Switch>
            <Route exact path='/' component={MainForm} />
            <Route path='/notFound' component={NotFound} />
            <Route path='/errorPage' component={ErrorPage} />
            <Route path='/:id' component={DisplaySecret} />
          </Switch>
        </Layout>
      </div>
    </BrowserRouter>
  );
}

export default App;
