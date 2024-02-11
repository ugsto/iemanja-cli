# Iemanja CLI

![Go](https://img.shields.io/badge/Go-00ADD8.svg?style=flat&logo=Go&logoColor=white)
[![tests](https://github.com/ugsto/iemanja-cli/actions/workflows/tests.yml/badge.svg)](https://github.com/ugsto/iemanja-cli/actions/workflows/tests.yml)

---

> [!CAUTION]
> This application is **not** production ready.

> [!WARNING]
> This application requires [`iemanjad`](https://github.com/ugsto/iemanjad) to be running.

Iemanja CLI is a command-line interface for managing posts and tags within the Iemanja application ecosystem. It allows users to perform CRUD operations on posts and tags through simple and intuitive commands.

## Installation

To install Iemanja CLI, use the following Go command:

```bash
go install github.com/ugsto/iemanja-cli/...@latest
```

## Usage

Iemanja CLI supports a variety of commands for managing posts and tags. Below are the available commands and their revised usage based on the updated CLI application code.

### Posts

- **List Posts**: Lists all posts with pagination support.
  ```bash
  iemanja-cli list-posts --limit 10 --offset 0
  ```
- **Create Post**: Creates a new post with a title, content, and tags.
  ```bash
  iemanja-cli create-post --title "Your Title" --content "Your Content" --tags tag1,tag2
  ```
- **Get Post**: Retrieves a post by its ID.
  ```bash
  iemanja-cli get-post <postID>
  ```
- **Update Post**: Updates an existing post's title, content, and tags by its ID.
  ```bash
  iemanja-cli update-post <postID> --title "New Title" --content "New Content" --tags newtag1,newtag2
  ```
- **Delete Post**: Deletes a post by its ID.
  ```bash
  iemanja-cli delete-post <postID>
  ```
- **Write Post**: Writes a new post or edits an existing post using the default editor.
  ```bash
  iemanja-cli write-post [<id>]
  ```

### Tags

- **List Tags**: Lists all tags with pagination support.
  ```bash
  iemanja-cli list-tags --limit 10 --offset 0
  ```
- **Create Tag**: Creates a new tag with a name.
  ```bash
  iemanja-cli create-tag <tagName>
  ```
- **Get Tag**: Retrieves a tag by its name.
  ```bash
  iemanja-cli get-tag <tagName>
  ```
- **Update Tag**: Updates an existing tag's name to a new name.
  ```bash
  iemanja-cli update-tag <originalName> --name "newName"
  ```
- **Delete Tag**: Deletes a tag by its name.
  ```bash
  iemanja-cli delete-tag <tagName>
  ```

## License

This project is licensed under the GNU Affero General Public License Version 3 (AGPLv3) - see the [LICENSE](LICENSE) file for details.
