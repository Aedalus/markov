# Markov

Markov is a text generation CLI tool written in go.

## Background

A [markov chain](https://en.wikipedia.org/wiki/Markov_chain) is a mathematical model describing a potential sequence of events. Think of a flow chart, if everytime you got to a fork in the road each path was weighted. The dice are rolled, and one of the paths are chosen at random. Each decision is also independent of anything but the current state, with no memory of any previous decisions.

When applied to text, this can be used to generate entire sentences. They often aren't fully meaningful, but manage to fake a decent grammatical structure.

### Examples generated from the works of HP Lovecraft

- People say the ice desert of bodies.
- March there came over what of that have we mustn't hesitate to mutter many declared it the bowsprit of visibility.
- Another moment of head, dragon body, prodigious claws.
- Wilbur had caused much so bad tonight.
- On November 26, 1916, proved to inhere in language reach them the noisome boarded place; and theosophist, and time two sections, the queer clay bas-relief, and traditions tell - as far down over these saxifrages.
- There came Burning burning IV.

## Usage

The first step is generating a TextDistribution from a given set of texts. This shows a breakdown of the frequencies between words.

```
// You can read in a single file at a time
markov dist "my-story.txt" -o my-story.json
// Or read in multiple texts at once, combining them
markov dist ./texts/**/*.txt -o all.json
```

From there, you are able to generate sentences from an output file.

```
// Read in the file and print out a sentence.
markov gen-sentence my-story.json
// Read in multiple files at once if you want.
markov gen-sentence my-story all.json
```
