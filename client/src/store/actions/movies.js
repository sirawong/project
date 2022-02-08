import { GET_MOVIES, SELECT_MOVIE } from '../types';
import { setAlert } from './alert';

export const uploadMovieImage = (id, image) => async dispatch => {
  try {
    const data = new FormData();
    console.log("test")
    data.append('file', image);
    const url = process.env.REACT_APP_BASE_MOVIE_URL + '/photo/' + id;
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

export const getMovies = () => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_MOVIE_URL ;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' }
    });
    const movies = await response.json();
    if (response.ok) {
      dispatch({ type: GET_MOVIES, payload: movies });
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

export const onSelectMovie = movie => ({
  type: SELECT_MOVIE,
  payload: movie
});

export const getMovie = id => async dispatch => {
  try {
    const url = process.env.REACT_APP_BASE_MOVIE_URL + '/' + id;
    const response = await fetch(url, {
      method: 'GET',
      headers: { 'Content-Type': 'application/json' }
    });
    const movie = await response.json();
    if (response.ok) {
      dispatch({ type: SELECT_MOVIE, payload: movie });
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

// export const getMovieSuggestion = id => async dispatch => {
//   try {
//     const url = process.env.REACT_APP_BASE_MOVIE_URL + '/usermodeling/' + id;
//     const response = await fetch(url, {
//       method: 'GET',
//       headers: { 'Content-Type': 'application/json' }
//     });
//     const movies = await response.json();
//     if (response.ok) {
//       dispatch({ type: GET_SUGGESTIONS, payload: movies });
//     }
//   } catch (error) {
//     dispatch(setAlert(error.message, 'error', 2000));
//   }
// };

export const addMovie = (image, newMovie) => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_MOVIE_URL;
    const response = await fetch(url, {
      method: 'POST',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "title":newMovie.title,
        "language":newMovie.language,
        "genre":newMovie.genre,
        "director":newMovie.director,
        "cast":newMovie.cast,
        "description":newMovie.description,
        "duration":String(newMovie.duration),
        "releaseDate":new Date(newMovie.releaseDate).toISOString(),
        "endDate":new Date(newMovie.endDate).toISOString(),
      })
    });
    const movie = await response.json();
    if (response.ok) {
      dispatch(setAlert('Movie have been saved!', 'success', 2000));
      if (image) dispatch(uploadMovieImage(movie._id, image));
      dispatch(getMovies());
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

export const updateMovie = (movieId, movie, image) => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_MOVIE_URL + '/' + movieId;
    const response = await fetch(url, {
      method: 'PUT',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        "title":movie.title,
        "language":movie.language,
        "genre":movie.genre,
        "director":movie.director,
        "cast":movie.cast,
        "description":movie.description,
        "duration":String(movie.duration),
        "releaseDate":new Date(movie.releaseDate).toISOString(),
        "endDate":new Date(movie.endDate).toISOString(),
      })
    });
    if (response.ok) {
      dispatch(onSelectMovie(null));
      dispatch(setAlert('Movie have been saved!', 'success', 2000));
      if (image) dispatch(uploadMovieImage(movieId, image));
      dispatch(getMovies());
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};

export const removeMovie = movieId => async dispatch => {
  try {
    const token = localStorage.getItem('jwtToken');
    const url = process.env.REACT_APP_BASE_MOVIE_URL + '/' + movieId;
    const response = await fetch(url, {
      method: 'DELETE',
      headers: {
        Authorization: `Bearer ${token}`,
        'Content-Type': 'application/json'
      }
    });
    if (response.ok) {
      dispatch(getMovies());
      dispatch(onSelectMovie(null));
      dispatch(setAlert('Movie have been Deleted!', 'success', 2000));
    }
  } catch (error) {
    dispatch(setAlert(error.message, 'error', 2000));
  }
};
