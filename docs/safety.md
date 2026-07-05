# Safety model

FotoForge never automatically deletes user data.

Read-only commands may catalog observations. A command that changes files must first produce a stable plan containing inputs, expected hashes, destinations, conflicts, and rollback information. Execution requires explicit user approval of that plan. Every step is journaled before the next begins, and verification follows execution.

Permanent removal, when implemented, will operate only on explicitly selected quarantine entries. It will use separate confirmation, refuse ambiguous or stale plans, and clearly state that the final step is irreversible.

Interrupted operations must be safe to resume or restore. Errors must identify the affected path and operation without leaking unrelated metadata. Tests for mutating workflows must cover partial completion, destination conflicts, changed inputs, permission failures, and restoration.
