import type { TMovie } from "../types";
import WatchedMovie from "./WatchedMovie";

type WatchedMovieListProps = {
  watched: TMovie[];
};

const WatchedMovieList = (props: WatchedMovieListProps) => {
  return (
    <ul className="list">
      {props.watched.map((movie) => (
        <WatchedMovie watched={movie} key={movie.imdbID} />
      ))}
    </ul>
  );
};

export default WatchedMovieList;
