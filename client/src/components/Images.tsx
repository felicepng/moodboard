import type { Component } from 'solid-js';
import { For } from 'solid-js';
import { images } from '../App';

const Images: Component = () => {
  return (
    <div class="grid grid-cols-2 md:grid-cols-4 gap-8">
      <For each={images().urls}>
        {(url: string) => (
          <img src={url} class="w-32 h-32 md:w-40 md:h-40 2xl:w-60 2xl:h-60" />
        )}
      </For>
    </div>
  );
};

export default Images;
