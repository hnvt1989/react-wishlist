import _ from 'lodash';
import {
  FETCH_WISH,
  FETCH_WISHES,
  CREATE_WISH,
  EDIT_WISH,
  DELETE_WISH,
} from '../actions/types';

export default (state = {}, action) => {
  switch (action.type) {
    case FETCH_WISHES:
      return { ...state, ..._.mapKeys(action.payload, 'id') };
    case FETCH_WISH:
      return { ...state, [action.payload.id]: action.payload };
    case CREATE_WISH:
      return { ...state, [action.payload.id]: action.payload };
    case EDIT_WISH:
      return { ...state, [action.payload.id]: action.payload };
    case DELETE_WISH:
      return _.omit(state, action.payload);
    default:
      return state;
  }
};
