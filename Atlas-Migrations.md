Sure! Here’s a ready-to-use Markdown file you can save as `Atlas-Migrations.md` for your project:

````markdown
# Atlas Migration Commands - Go-SaaS-Kit

> Make sure your `DATABASE_URL` environment variable is set before running these commands:

```bash
export DATABASE_URL="postgresql://user:password@host:port/dbname?sslmode=require&search_path=public"
```
````

---

## Migration Commands

| Command                | Description                                                                                       |
| ---------------------- | ------------------------------------------------------------------------------------------------- |
| `make atlas-diff`      | Generate a new migration from your GORM models. Prompts for a name; uses timestamp if left empty. |
| `make atlas-apply-dry` | Preview SQL statements **without applying** changes to the database.                              |
| `make atlas-lint`      | Validate migration safety and check for unsafe operations.                                        |
| `make atlas-apply`     | Apply migrations to the database and update `atlas_schema_revisions`.                             |

---

## Recommended Workflow

1. Update your GORM models.
2. Generate a migration:

```bash
make atlas-diff
```

3. Preview migration SQL:

```bash
make atlas-apply-dry
```

4. Check safety:

```bash
make atlas-lint
```

5. Apply migration:

```bash
make atlas-apply
```

> This workflow ensures your database schema stays in sync with your models safely.

```

---

```
