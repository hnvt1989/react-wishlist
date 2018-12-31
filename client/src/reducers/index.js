import { combineReducers } from 'redux';
import { reducer as formReducer } from 'redux-form';
import authReducer from './authReducer';
import wishReducer from './wishReducer';
import currentUserReducer from './currentUserReducer';

export default combineReducers({
  auth: authReducer,
  form: formReducer,
  wishes: wishReducer,
  currentUser: currentUserReducer
});
