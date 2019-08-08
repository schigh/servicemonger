package ui

const imgB64 = `iVBORw0KGgoAAAANSUhEUgAAACQAAAAkCAYAAADhAJiYAAAEJWlUWHRYTUw6Y29tLmFkb2JlLnhtcAAAAAAAPD94cGFja2V0IGJlZ2luPSLvu78iIGlkPSJXNU0wTXBDZWhpSHpyZVN6TlRjemtjOWQiPz4KPHg6eG1wbWV0YSB4bWxuczp4PSJhZG9iZTpuczptZXRhLyIgeDp4bXB0az0iWE1QIENvcmUgNS41LjAiPgogPHJkZjpSREYgeG1sbnM6cmRmPSJodHRwOi8vd3d3LnczLm9yZy8xOTk5LzAyLzIyLXJkZi1zeW50YXgtbnMjIj4KICA8cmRmOkRlc2NyaXB0aW9uIHJkZjphYm91dD0iIgogICAgeG1sbnM6ZXhpZj0iaHR0cDovL25zLmFkb2JlLmNvbS9leGlmLzEuMC8iCiAgICB4bWxuczp0aWZmPSJodHRwOi8vbnMuYWRvYmUuY29tL3RpZmYvMS4wLyIKICAgIHhtbG5zOnhtcD0iaHR0cDovL25zLmFkb2JlLmNvbS94YXAvMS4wLyIKICAgIHhtbG5zOnhtcE1NPSJodHRwOi8vbnMuYWRvYmUuY29tL3hhcC8xLjAvbW0vIgogICAgeG1sbnM6c3RFdnQ9Imh0dHA6Ly9ucy5hZG9iZS5jb20veGFwLzEuMC9zVHlwZS9SZXNvdXJjZUV2ZW50IyIKICAgZXhpZjpQaXhlbFhEaW1lbnNpb249IjM2IgogICBleGlmOlBpeGVsWURpbWVuc2lvbj0iMzYiCiAgIHRpZmY6SW1hZ2VXaWR0aD0iMzYiCiAgIHRpZmY6SW1hZ2VMZW5ndGg9IjM2IgogICB0aWZmOlJlc29sdXRpb25Vbml0PSIyIgogICB0aWZmOlhSZXNvbHV0aW9uPSIxNDQuMCIKICAgdGlmZjpZUmVzb2x1dGlvbj0iMTQ0LjAiCiAgIHhtcDpNb2RpZnlEYXRlPSIyMDE5LTA4LTA2VDE0OjM2OjIzLTA0OjAwIgogICB4bXA6TWV0YWRhdGFEYXRlPSIyMDE5LTA4LTA2VDE0OjM2OjIzLTA0OjAwIj4KICAgPHhtcE1NOkhpc3Rvcnk+CiAgICA8cmRmOlNlcT4KICAgICA8cmRmOmxpCiAgICAgIHN0RXZ0OmFjdGlvbj0icHJvZHVjZWQiCiAgICAg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`