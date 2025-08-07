import type { TMovie } from "../types";
import WatchedMovie from "./WatchedMovie";

type WatchedMovieListProps = {
  onDelete: (id: string) => void;
  watched: TMovie[];
};

const WatchedMovieList = ({ watched, onDelete }: WatchedMovieListProps) => {
  return (
    <ul className="list">
      {watched.map((movie) => (
        <WatchedMovie watched={movie} key={movie.imdbID} onDelete={onDelete} />
      ))}
    </ul>
  );
};

export default WatchedMovieList;
