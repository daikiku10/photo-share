// Code generated by "accessor -type=Photo"; DO NOT EDIT.

package photo

import "time"

// Id return id value
func (t Photo) Id() Id {
	return t.id
}

// Title return title value
func (t Photo) Title() string {
	return t.title
}

// Description return description value
func (t Photo) Description() string {
	return t.description
}

// ImageUrl return imageUrl value
func (t Photo) ImageUrl() string {
	return t.imageUrl
}

// AuthorId return authorId value
func (t Photo) AuthorId() string {
	return t.authorId
}

// CategoryId return categoryId value
func (t Photo) CategoryId() string {
	return t.categoryId
}

// CreatedAt return createdAt value
func (t Photo) CreatedAt() time.Time {
	return t.createdAt
}

// LikedCount return likedCount value
func (t Photo) LikedCount() int {
	return t.likedCount
}
