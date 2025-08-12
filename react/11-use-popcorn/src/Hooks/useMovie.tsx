import { useEffect, useState } from "react";
import type { TMovie } from "../types";

export const key = "644cebbb";

export function useMovies(query: string) {
  const [movies, setMovies] = useState<TMovie[]>([]);
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [err, setError] = useState<string>("");

  useEffect(
    function () {
      async function fetchMovies() {
        try {
          const queryUrl = `http://www.omdbapi.com/?apikey=${key}&s=${query}`;
          setIsLoading(true);
          setError("");
          const res = await fetch(queryUrl);

          if (!res.ok)
            throw new Error("Something went wrong while fetching movies");

          const data = await res.json();
          if (data.Response === "False") {
            throw new Error("Movie not found");
          }

          const movies: TMovie[] = data.Search.map((movie: any) => ({
            imdbID: movie.imdbID,
            title: movie.Title,
            year: movie.Year,
            poster: movie.Poster,
          }));

          setMovies(movies);
          setError("");
        } catch (err) {
          if (err instanceof Error) {
            console.log(err.message);
            if (err.name !== "AbortError") {
              setError(err.message);
            }
          }
        } finally {
          setIsLoading(false);
        }
      }

      if (query.length <= 3) {
        setMovies([]);
        setError("");
        return;
      }

      // handleCloseMovie()
      fetchMovies();
    },
    [query],
  );

  return { movies, isLoading, err };
}
