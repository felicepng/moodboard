import type { Component } from 'solid-js';
import { For } from 'solid-js';

interface Props {
  prompts: string[];
}

const Prompts: Component<Props> = (props: Props) => {
  const { prompts } = props;

  return (
    <div class="flex flex-col gap-y-4 items-center text-center text-lg md:text-2xl text-slate-400">
      <p>generated image prompts:</p>
      <div class="grid md:grid-cols-2 md:gap-x-20">
        <For each={prompts}>
          {(prompt, idx) => (
            <p>
              {idx() + 1}. {prompt}
            </p>
          )}
        </For>
      </div>
    </div>
  );
};

export default Prompts;
