import {
  REGISTER_SUCCESS,
  REGISTER_FAIL,
  USER_LOADED,
  AUTH_ERROR,
  LOGIN_SUCCESS,
  LOGIN_FAIL,
  LOGOUT
} from '../types';
import { setAlert } from './alert';
import { setAuthHeaders, setUser, removeUser, isLoggedIn } from '../../utils';

export const uploadImage = (id, image) => async dispatch => {
  try {
    const data = new FormData();
    data.append('file', image);
    const url = process.env.REACT_APP_BASE_USER_URL + '/photo/' + id;
    const response = await fetch(url, {
      method: 'POST',
      body: data
    });
    console.log("test")
    const responseData = await response.json();
    if (response.ok) {
      dispatch(setAlert('Image Uploaded', 'success', 2000));
    }
    if (responseData.error) {
      dispatch(setAlert(responseData.error.message, 'error', 2000));
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

// Login user
export const login = (username, password) => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_USER_URL + '/login';
    const response = await fetch(url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({ username, password })
    });
    const responseData = await response.json();
    if (response.ok) {
      const { user } = responseData;
      user && setUser(user);
      dispatch({ type: LOGIN_SUCCESS, payload: responseData });
      dispatch(setAlert(`Welcome ${user.name}`, 'success', 2000));
    }
    if (responseData.error) {
      dispatch({ type: LOGIN_FAIL });
      dispatch(setAlert(responseData.error.message, 'error', 2000));
    }
  } catch (error) {
    dispatch({ type: LOGIN_FAIL });
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

// export const facebookLogin = e => async dispatch => {
//   try {
//     const { email, userID, name } = e;
//     const options = {
//       method: 'POST',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify({ email, userID, name })
//     };
//     const url = process.env.REACT_APP_BASE_USER_URL + 'login/facebook';
//     const response = await fetch(url, options);
//     const responseData = await response.json();

//     if (response.ok) {
//       const { user } = responseData;
//       user && setUser(user);
//       dispatch({ type: LOGIN_SUCCESS, payload: responseData });
//       dispatch(setAlert(`Welcome ${user.name}`, 'success', 2000));
//     }
//     if (responseData.error) {
//       dispatch({ type: LOGIN_FAIL });
//       dispatch(setAlert(responseData.error.message, 'error', 2000));
//     }
//   } catch (error) {
//     dispatch({ type: LOGIN_FAIL });
//     dispatch(setAlert(error.message, 'error', 2000));
//   }
// };

// export const googleLogin = ({ profileObj }) => async dispatch => {
//   try {
//     const { email, googleId, name } = profileObj;
//     const options = {
//       method: 'POST',
//       headers: { 'Content-Type': 'application/json' },
//       body: JSON.stringify({ email, googleId, name })
//     };
//     const url = process.env.REACT_APP_BASE_USER_URL + 'login/google';
//     const response = await fetch(url, options);
//     const responseData = await response.json();

//     if (response.ok) {
//       const { user } = responseData;
//       user && setUser(user);
//       dispatch({ type: LOGIN_SUCCESS, payload: responseData });
//       dispatch(setAlert(`Welcome ${user.name}`, 'success', 2000));
//     }
//     if (responseData.error) {
//       dispatch({ type: LOGIN_FAIL });
//       dispatch(setAlert(responseData.error.message, 'error', 2000));
//     }
//   } catch (error) {
//     dispatch({ type: LOGIN_FAIL });
//     dispatch(setAlert(error.message, 'error', 2000));
//   }
// };

// Register user
export const register = ({
  name,
  username,
  email,
  phone,
  image,
  password
}) => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_USER_URL;
    const body = { name, username, email, phone, password };
    const response = await fetch(url, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(body)
    });
    const responseData = await response.json();
    if (response.ok) {
      const { user } = responseData;
      user && setUser(user);
      if (image) dispatch(uploadImage(user._id, image)); // Upload image
      dispatch({ type: REGISTER_SUCCESS, payload: responseData });
      dispatch(setAlert('Register Success', 'success', 2000));
    }
    if (responseData._message) {
      dispatch({ type: REGISTER_FAIL });
      dispatch(setAlert(responseData.message, 'error', 2000));
    }
  } catch (error) {
    dispatch({ type: REGISTER_FAIL });
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

// Load user
export const loadUser = () => async dispatch => {
  if (!isLoggedIn()) return;
  try {
    const url = process.env.REACT_APP_BASE_USER_URL + '/me';
    const response = await fetch(url, {
      method: 'GET',
      headers: setAuthHeaders()
    });
    const responseData = await response.json();
    if (response.ok) {
      const { user } = responseData;
      user && setUser(user);
      dispatch({ type: USER_LOADED, payload: responseData });
    }
    if (!response.ok) dispatch({ type: AUTH_ERROR });
  } catch (error) {
    dispatch({ type: AUTH_ERROR });
  }
};

// Logout
export const logout = () => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url =   process.env.REACT_APP_BASE_USER_URL + '/logout';
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    });
    const responseData = await response.json();
    if (response.ok) {
      removeUser();
      dispatch({ type: LOGOUT });
      dispatch(setAlert('LOGOUT Success', 'success', 2000));
    }
    if (responseData.error) {
      dispatch(setAlert(responseData.error.message, 'error', 2000));
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};
