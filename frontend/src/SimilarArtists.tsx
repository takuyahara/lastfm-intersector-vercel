import { createSignal, createContext, useContext, JSX } from "solid-js";

export type Props = {
  value?: Value;
  children: JSX.Element;
};
export type Value = [
  arg0: () => SimilarArtists[],
  arg1: (arg0: SimilarArtists[]) => void
];
export interface Artist {
  name: string;
  match: number;
  url: string;
}

export type SimilarArtists = {
  artist: string;
  similarartists: Artist[];
};

export type SimilarArtistsWithError = {
  error: string;
} & SimilarArtists;

const SimilarArtistsContext = createContext<Value>();

export function SimilarArtistsProvider(props: Props) {
  const [similarArtists, setSimilarArtists] = createSignal<SimilarArtists[]>(
    []
  );
  const value: Value = [similarArtists, setSimilarArtists];
  return (
    <SimilarArtistsContext.Provider value={value}>
      {props.children}
    </SimilarArtistsContext.Provider>
  );
}

export function useSimilarArtists() {
  return useContext(SimilarArtistsContext);
}
