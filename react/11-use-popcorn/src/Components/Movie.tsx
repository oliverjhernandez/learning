import type { TMovie } from "../types";

type MovieProps = {
  onSelectMovie: (id: string) => void;
  movie: TMovie;
};

const Movie = ({ onSelectMovie, movie }: MovieProps) => {
  return (
    <li
      key={movie.imdbID}
      onClick={() => {
        onSelectMovie(movie.imdbID);
      }}
    >
      <img src={movie.poster} alt={`${movie.title} poster`} />
      <h3>{movie.title}</h3>
      <div>
        <p>
          <span>🗓</span>
          <span>{movie.year}</span>
        </p>
      </div>
    </li>
  );
};

export default Movie;
