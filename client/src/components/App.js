import React from 'react';
import { Router, Route, Switch } from 'react-router-dom';
import WishCreate from './wishes/WishCreate';
import WishEdit from './wishes/WishEdit';
import WishDelete from './wishes/WishDelete';
import WishList from './wishes/WishList';
import WishShow from './wishes/WishShow';
import Header from './Header';
import history from '../history';

const App = () => {
  return (
    <div className="ui container">
      <Router history={history}>
        <div>
          <Header />
          <Switch>
            <Route path="/" exact component={WishList} />
            <Route path="/wishes/new" exact component={WishCreate} />
            <Route path="/wishes/edit/:id" exact component={WishEdit} />
            <Route path="/wishes/delete/:id" exact component={WishDelete} />
            <Route path="/wishes/:id" exact component={WishShow} />
          </Switch>
        </div>
      </Router>
    </div>
  );
};

export default App;
