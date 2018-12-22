import wishes from '../apis/wishes';
import history from '../history';
import {
  SIGN_IN,
  SIGN_OUT,
  CREATE_WISH,
  FETCH_WISHES,
  FETCH_WISH,
  DELETE_WISH,
  EDIT_WISH
} from './types';

export const signIn = userId => {
  return {
    type: SIGN_IN,
    payload: userId
  };
};

export const signOut = () => {
  return {
    type: SIGN_OUT
  };
};

export const createWish = formValues => async (dispatch, getState) => {
  const { userId } = getState().auth;
  const response = await wishes.post('/wishes', { ...formValues, userId });

  dispatch({ type: CREATE_WISH, payload: response.data });
  history.push('/');
};

export const fetchWishes = () => async dispatch => {
  const response = await wishes.get('/wishes');

  dispatch({ type: FETCH_WISHES, payload: response.data });
};

export const fetchWish= id => async dispatch => {
  const response = await wishes.get(`/wishes/${id}`);

  dispatch({ type: FETCH_WISH, payload: response.data });
};

export const editWish = (id, formValues) => async dispatch => {
  const response = await wishes.patch(`/wishes/${id}`, formValues);

  dispatch({ type: EDIT_WISH, payload: response.data });
  history.push('/');
};

export const deleteWish = id => async dispatch => {
  await wishes.delete(`/wishes/${id}`);

  dispatch({ type: DELETE_WISH, payload: id });
  history.push('/');
};
