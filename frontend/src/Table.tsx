import { JSX } from "solid-js";
import { sprintf } from "sprintf-js";
import { Artist } from "./SimilarArtists";

interface FormProps {
  artist: string;
  similarartists: Artist[];
}

export default function Table(props: FormProps): JSX.Element {
  const similarArtistExists = props.similarartists.length > 0;
  return (
    <div
      style="height: min-content"
      class="w-full h-min border(gray-500 1) divide-y divide-gray-500 rounded-2xl shadow-lg"
    >
      <strong
        class={`block px-5 py-3 text-xl text-gray-200 bg-gray-700 ${
          similarArtistExists ? "rounded-t-2xl" : "rounded-2xl"
        }`}
      >
        {props.artist}
      </strong>
      {similarArtistExists && (
        <ul class="divide-y divide-gray-500">
          {props.similarartists.map((sa) => (
            <li class="last:rounded-b-2xl bg-gray-700 hover:bg-gray-600">
              <a
                class="px-5 py-3 block"
                href={sa.url}
                target="_blank"
                rel="noopener noreferrer"
              >
                <strong class="block mb-1 text-gray-300">{sa.name}</strong>
                <div class="text-gray-400 text-xs">
                  Match: {sprintf("%.6f", sa.match)}
                </div>
              </a>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}
