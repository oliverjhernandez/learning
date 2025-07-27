import type { TMovie } from "../types";

type WatchedMovieProps = {
  watched: TMovie;
};

const WatchedMovie = (props: WatchedMovieProps) => {
  return (
    <li>
      <img src={props.watched.Poster} alt={`${props.watched.Title} poster`} />
      <h3>{props.watched.Title}</h3>
      <div>
        <p>
          <span>⭐️</span>
          <span>{props.watched.imdbRating}</span>
        </p>
        <p>
          <span>🌟</span>
          <span>{props.watched.userRating}</span>
        </p>
        <p>
          <span>⏳</span>
          <span>{props.watched.runtime} min</span>
        </p>
      </div>
    </li>
  );
};

export default WatchedMovie;
