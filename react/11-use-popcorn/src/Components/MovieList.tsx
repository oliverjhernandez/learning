import type { TMovie } from "../types";
import Movie from "./Movie";

type MovieListProps = {
  onSelectMovie: (id: string) => void;
  movies: TMovie[];
};

const MovieList = ({ onSelectMovie, movies }: MovieListProps) => {
  return (
    <ul className="list list-movies">
      {movies.map((movie: TMovie) => (
        <Movie key={movie.imdbID} movie={movie} onSelectMovie={onSelectMovie} />
      ))}
    </ul>
  );
};

export default MovieList;
