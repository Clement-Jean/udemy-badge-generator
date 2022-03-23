# udemy-badge-generator

This action creates a udemy rating badge for one or multiple courses (average).

example: ![Udemy](.github/badges/udemy.svg)

## Environment Variables

**It is recommended to set these as Github Secrets and use the secrets inside your github action yml file (see: https://docs.github.com/en/actions/security-guides/encrypted-secrets)**

### `UDEMY_TOKEN`

**Required** Token generated for call to instructor API (see: https://www.udemy.com/instructor/account/api/)

### `UDEMY_COURSE`

**Required** One or multiple course ID separated by commas (see: https://www.udemy.com/developers/instructor/methods/get-api-taught-courses-list/)

## Inputs

## `label`

**Optional** Label text. Default `"udemy"`.

### `label-bg`

**Optional** Label background color. Default `"#983DE7"` (udemy logo color).

### `label-fg`

**Optional** Label foreground color. Default `"#FFF"`.

### `rating-fg`

**Optional** Rating foreground color. Default `"#FFF"`.

### `badge-radius`

**Optional** Radius for rounding corners of the badge. Default `3`

### `generate-path`

**Optional** Path where the badge will be generated. Default `./github/badges`

### `generate-filename`

**Optional** Filename, without extension, of the file generated. Default `udemy`

## Example usage

```yml
uses: Clement-Jean/udemy-badge-generator@v1
env:
  UDEMY_TOKEN: ${{ secrets.UDEMY_TOKEN }}
  UDEMY_COURSE: ${{ secrets.UDEMY_COURSE }}
```
