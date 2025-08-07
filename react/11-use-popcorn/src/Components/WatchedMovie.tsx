import type { TMovie } from "../types";

type WatchedMovieProps = {
  onDelete: (id: string) => void;
  watched: TMovie;
};

const WatchedMovie = ({ watched, onDelete }: WatchedMovieProps) => {
  return (
    <li>
      <img src={watched.poster} alt={`${watched.title} poster`} />
      <h3>{watched.title}</h3>
      <div>
        <p>
          <span>⭐️</span>
          <span>{watched.imdbRating}</span>
        </p>
        <p>
          <span>🌟</span>
          <span>{watched.userRating}</span>
        </p>
        <p>
          <span>⏳</span>
          <span>{watched.runtime} min</span>
        </p>
        <button className="btn-delete" onClick={() => onDelete(watched.imdbID)}>
          X
        </button>
      </div>
    </li>
  );
};

export default WatchedMovie;
