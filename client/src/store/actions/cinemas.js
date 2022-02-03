import { GET_CINEMAS, GET_CINEMA } from '../types';
import { setAlert } from './alert';

export const uploadCinemaImage = (id, image) => async dispatch => {
  try {
    const data = new FormData();
    data.append('file', image);
    const url = process.env.REACT_APP_BASE_CINEMA_URL + '/photo/' + id;
    const response = await fetch(url, {
      method: 'POST',
      body: data
    });
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

export const getCinemas = () => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_CINEMA_URL;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' }
    });
    const cinemas = await response.json();
    if (response.ok) {
      dispatch({ type: GET_CINEMAS, payload: cinemas });
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

export const getCinema = id => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_CINEMA_URL + '/' + id;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' }
    });
    const cinema = await response.json();
    if (response.ok) {
      dispatch({ type: GET_CINEMA, payload: cinema });
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

export const createCinemas = (image, newCinema) => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_CINEMA_URL;
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "name": newCinema.name,
        "ticketPrice": Number(newCinema.ticketPrice),
        "city": newCinema.city,
        "seats": newCinema.seats,
        "seatsAvailable": Number(newCinema.seatsAvailable),
      })
    });
    const cinema = await response.json();
    if (response.ok) {
      dispatch(setAlert('Cinema Created', 'success', 2000));
      if (image) dispatch(uploadCinemaImage(cinema._id, image));
      dispatch(getCinemas());
      return { status: 'success', message: 'Cinema Created' };
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
    return {
      status: 'error',
      message: ' Cinema have not been saved, try again.'
    };
  }
};

export const updateCinemas = (image, cinema, id) => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_CINEMA_URL + '/'  + id;
    const response = await fetch(url, {
      method: 'PATCH',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "name": cinema.name,
        "ticketPrice": Number(cinema.ticketPrice),
        "city": cinema.city,
        "seats": cinema.seats,
        "seatsAvailable": Number(cinema.seatsAvailable),
      })
    });
    if (response.ok) {
      dispatch(setAlert('Cinema Updated', 'success', 2000));
      if (image) dispatch(uploadCinemaImage(id, image));
      return { status: 'success', message: 'Cinema Updated' };
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
    return {
      status: 'error',
      message: ' Cinema have not been updated, try again.'
    };
  }
};

export const removeCinemas = id => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_CINEMA_URL + '/'  + id;
    const response = await fetch(url, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    });
    if (response.ok) {
      dispatch(setAlert('Cinema Deleted', 'success', 2000));
      return { status: 'success', message: 'Cinema Removed' };
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
    return {
      status: 'error',
      message: ' Cinema have not been deleted, try again.'
    };
  }
};

// export const getCinemasUserModeling = username => async dispatch => {
//   try {
//     const url = process.env.REACT_APP_BASE_CINEMA_URL + '/usermodeling/' + username;
//     const response = await fetch(url, {
//       method: 'GET',
//       headers: { 'Content-Type': 'application/json' }
//     });
//     const cinemas = await response.json();
//     if (response.ok) {
//       dispatch({ type: GET_CINEMAS, payload: cinemas });
//     }
//   } catch (error) {
//     dispatch(setAlert(error.message, 'error', 2000));
//   }
// };
