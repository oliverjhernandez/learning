import type { TMovie } from "../types";
import Movie from "./Movie";

type MovieListProps = {
  movies: TMovie[];
};

const MovieList = (props: MovieListProps) => {
  return (
    <ul className="list">
      {props.movies.map((movie: TMovie) => (
        <Movie movie={movie} key={movie.imdbID} />
      ))}
    </ul>
  );
};

export default MovieList;
