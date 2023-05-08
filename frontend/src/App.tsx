import type { Component } from "solid-js";
import Form from "./Form";
import Table from "./Table";
import { useSimilarArtists } from "./SimilarArtists";
import { IntersectedArtists } from "./IntersectedArtists";

const App: Component = () => {
  const [similarArtists] = useSimilarArtists();
  return (
    <div class="min-h-screen bg-gray-600">
      <header class="px-10 py-8 w-full bg-gray-700 shadow-lg">
        <Form></Form>
      </header>
      <main class="h-120 flex gap-7 w-full h-full px-10 py-10">
        {similarArtists().length > 0 && (
          <IntersectedArtists similarArtists={similarArtists()} />
        )}
        {similarArtists().map((sims) => (
          <Table artist={sims.artist} similarartists={sims.similarartists} />
        ))}
      </main>
    </div>
  );
};

export default App;
