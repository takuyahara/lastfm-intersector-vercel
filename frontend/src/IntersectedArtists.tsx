import { JSX } from "solid-js";
import { Artist, SimilarArtists } from "./SimilarArtists";
import Table from "./Table";

const intersect = function (
  similarArtists: SimilarArtists[]
): Record<string, Artist[]> {
  const hash: Record<string, Artist[]> = {};
  for (const sim of [...similarArtists]) {
    for (const s of sim.similarartists) {
      if (hash[s.url] === undefined) {
        hash[s.url] = [];
      }
      hash[s.url].push(s);
    }
  }
  return hash;
};

const calculate = function (
  keys: number,
  record: Record<string, Artist[]>
): Artist[] {
  const recordsFiltered = Object.values(record).filter(
    (rec) => rec.length === keys
  );
  const calculated = recordsFiltered.reduce((acc, cur) => {
    const a = cur.slice(1).reduce((a, c) => {
      a.match *= c.match;
      return a;
    }, cur[0]);
    acc.push(a);
    return acc;
  }, [] as Artist[]);
  return calculated;
};

export type Props = {
  similarArtists: SimilarArtists[];
};

export function IntersectedArtists(props: Props): JSX.Element {
  const { similarArtists } = props;
  const step1 = intersect(similarArtists);
  const step2 = calculate(similarArtists.length, step1);
  step2.sort((a, b) => b.match - a.match);
  return <Table artist="Intersected" similarartists={step2} />;
}
