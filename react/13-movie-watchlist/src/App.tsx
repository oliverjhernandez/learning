import { useState, type ReactNode } from "react";
import "./index.css";

const initialMovies: Movie[] = [
  { id: 1, title: "The Matrix", genre: "Sci-Fi", year: 1999 },
  { id: 2, title: "Inception", genre: "Sci-Fi", year: 2010 },
  { id: 3, title: "The Shawshank Redemption", genre: "Drama", year: 1994 },
  { id: 4, title: "Pulp Fiction", genre: "Crime", year: 1994 },
  { id: 5, title: "Spirited Away", genre: "Animation", year: 2001 },
  { id: 6, title: "The Dark Knight", genre: "Action", year: 2008 },
  { id: 7, title: "La La Land", genre: "Musical", year: 2016 },
  { id: 8, title: "Parasite", genre: "Thriller", year: 2019 },
  { id: 9, title: "Forrest Gump", genre: "Drama", year: 1994 },
  { id: 10, title: "Interstellar", genre: "Sci-Fi", year: 2014 },
];

const initialToWatch: Movie[] = [];

type Movie = {
  id: number;
  title: string;
  genre: string;
  year: number;
};

const App = () => {
  const [movies, setMovies] = useState(initialMovies);
  const [moviesToWatch, setMoviesToWatch] = useState(initialToWatch);

  const handleToggleAddToList = (movie: Movie) => {
    if (moviesToWatch.some((m) => m.id === movie.id)) {
      setMoviesToWatch(moviesToWatch.filter((m) => m.id !== movie.id));
    } else {
      setMoviesToWatch([...moviesToWatch, movie]);
    }
  };

  return (
    <div className="app">
      <h1>Movie Watchlist</h1>
      <Section className="movie-list">
        <MovieList
          title="All Movies"
          movies={movies}
          onToggleAddToList={handleToggleAddToList}
          buttonLabel="Add to Watchlist"
        />
      </Section>
      <Section className="watchlist">
        <MovieList
          title="My Watchlist"
          movies={moviesToWatch}
          onToggleAddToList={handleToggleAddToList}
          buttonLabel="Remove from Watchlist"
        />
        {moviesToWatch.length < 1 && <p>No movies in watchlist</p>}
      </Section>
    </div>
  );
};

type MovieListProps = {
  movies: Movie[];
  title: string;
  onToggleAddToList: (m: Movie) => void;
  buttonLabel: string;
};

const MovieList = ({
  movies,
  title,
  onToggleAddToList,
  buttonLabel,
}: MovieListProps) => {
  return (
    <div>
      <h2>{title}</h2>
      {movies.map((m) => {
        return (
          <Movie
            key={m.id}
            movie={m}
            buttonLabel={buttonLabel}
            onToggleAddToList={onToggleAddToList}
          />
        );
      })}
    </div>
  );
};

type MovieProps = {
  movie: Movie;
  onToggleAddToList: (m: Movie) => void;
  buttonLabel: string;
};

const Movie = ({ movie, onToggleAddToList, buttonLabel }: MovieProps) => {
  return (
    <div className="movie-card">
      <h3>{movie.title}</h3>
      <p>
        {movie.genre} ({movie.year})
      </p>
      {onToggleAddToList && (
        <Button onClick={() => onToggleAddToList(movie)}>{buttonLabel}</Button>
      )}
    </div>
  );
};

type ButtonProps = {
  children: ReactNode;
  onClick: () => void;
};

const Button = ({ children, onClick }: ButtonProps) => {
  return <button onClick={onClick}>{children}</button>;
};

type SectionProps = {
  children: ReactNode;
  className: string;
};

const Section = ({ children, className }: SectionProps) => {
  return <section className={className}>{children}</section>;
};

export default App;
