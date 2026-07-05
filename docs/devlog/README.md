# FotoForge Development Log

Devlogs record meaningful development sessions so future contributors can understand what changed, why decisions were made, how the work was verified, and what remains unresolved. They complement commit history by preserving context that a diff cannot show.

## When to write an entry

Write a devlog after work that establishes or changes:

- product direction or architecture;
- safety constraints or archive-handling behavior;
- important implementation or workflow decisions;
- a substantial feature, migration, investigation, or incident response; or
- findings that will materially affect later work.

Routine typo fixes and mechanical maintenance generally do not need an entry. Agents should suggest an entry when a session qualifies, but must not invent outcomes or claim planned work was completed.

## Naming convention

Use:

```text
YYYY-MM-DD-short-topic.md
```

Keep the topic lowercase and separate words with hyphens. For example:

```text
2026-07-05-audit-design.md
```

Create a dated entry from the repository root with:

```sh
make devlog slug=audit-design
```

The helper uses the current system date and refuses to overwrite an existing entry.

## Entry contents

Start from [TEMPLATE.md](TEMPLATE.md). Each entry should capture:

- the date and a concise session summary;
- the problem or context that motivated the work;
- changes actually made;
- architectural, product, safety, or workflow decisions and their rationale;
- commands or evidence used to verify the result;
- remaining risks, limitations, and TODOs; and
- concrete next steps.

Keep entries factual, concise, and useful to someone returning months later. Clearly distinguish completed work from proposals, roadmap ideas, and unverified assumptions. Do not include credentials, personal media, private metadata, or sensitive filesystem details.
