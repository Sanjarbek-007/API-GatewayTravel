// models.go

package models

import "time"

type Author struct {
    ID       string `json:"id"`
    Username string `json:"username"`
    FullName string `json:"full_name"`
}

type Story struct {
    ID         string    `json:"id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    Location   string    `json:"location"`
    Tags       []string  `json:"tags"`
    AuthorID   string    `json:"author_id"`
    CreatedAt  time.Time `json:"created_at"`
}

type CrateStoryRequest struct {
    Title     string   `json:"title"`
    Content   string   `json:"content"`
    Location  string   `json:"location"`
    Tags      []string `json:"tags"`
    AuthorID  string   `json:"author_id"`
}

type CrateStoryResponse struct {
    ID         string    `json:"id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    Location   string    `json:"location"`
    Tags       []string  `json:"tags"`
    AuthorID   string    `json:"author_id"`
    CreatedAt  string    `json:"created_at"`
}

type UpdateStoryRequest struct {
    StoryID string `json:"story_id"`
    Title   string `json:"title"`
    Content string `json:"content"`
}

type UpdateStoryResponse struct {
    ID         string    `json:"id"`
    Title      string    `json:"title"`
    Content    string    `json:"content"`
    Location   string    `json:"location"`
    Tags       []string  `json:"tags"`
    Images     []string  `json:"images"`
    AuthorID   string    `json:"author_id"`
    UpdatedAt  string    `json:"updated_at"`
}

type DeleteStoryRequest struct {
    StoryID string `json:"story_id"`
}

type DeleteStoryResponse struct {
    MessageStory   bool   `json:"message_story"`
    Location       string `json:"location"`
    LikesCount     int32  `json:"likes_count"`
    CommentsCount  int32  `json:"comments_count"`
    CreatedAt      string `json:"created_at"`
}

type GetAllStoriesRequest struct {
    AuthorID string `json:"author_id"`
    Limit    int32  `json:"limit"`
    Offset   int32  `json:"offset"`
}

type GetAllStoriesResponse struct {
    Stories []Story `json:"stories"`
}

type StoryFullInfoRequest struct {
    StoryID string `json:"story_id"`
}

type StoryFullInfoResponse struct {
    ID            string    `json:"id"`
    Title         string    `json:"title"`
    Content       string    `json:"content"`
    Location      string    `json:"location"`
    Tags          []string  `json:"tags"`
    Author        Author    `json:"author"`
    LikesCount    int32     `json:"likes_count"`
    CommentsCount int32     `json:"comments_count"`
    CreatedAt     string    `json:"created_at"`
    UpdatedAt     string    `json:"updated_at"`
}

type CommentStoryRequest struct {
    StoryID  string `json:"story_id"`
    AuthorID string `json:"author_id"`
    Content  string `json:"content"`
}

type CommentStoryResponse struct {
    ID         string `json:"id"`
    Content    string `json:"content"`
    AuthorID   string `json:"author_id"`
    StoryID    string `json:"story_id"`
    CreatedAt  string `json:"created_at"`
}

type Comment struct {
    ID         string    `json:"id"`
    Content    string    `json:"content"`
    Author     Author    `json:"author"`
    CreatedAt  string    `json:"created_at"`
}

type GetAllCommentRequest struct {
    StoryID string `json:"story_id"`
    Limit   int32  `json:"limit"`
    Offset  int32  `json:"offset"`
}

type GetAllCommentResponse struct {
    Comments  []Comment `json:"comments"`
    CreatedAt string    `json:"created_at"`
}

type CreateLikeRequest struct {
    StoryID string `json:"story_id"`
    UserID  string `json:"user_id"`
}

type CreateLikeResponse struct {
    StoryID string `json:"story_id"`
    UserID  string `json:"user_id"`
    LikedAt string `json:"liked_at"`
}

type Destination struct {
    Name              string   `json:"name"`
    StartDate         string   `json:"start_date"`
    EndDate           string   `json:"end_date"`
    PopularActivities []string `json:"popular_activities"`
}

type CreateItinerariesRequest struct {
    Title        string      `json:"title"`
    Description  string      `json:"description"`
    StartDate    string      `json:"start_date"`
    EndDate      string      `json:"end_date"`
    Destinations Destination `json:"destinations"`
}

type CreateItinerariesResponse struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    AuthorID   string `json:"author_id"`
    CreatedAt  string `json:"created_at"`
}

type UpdateItinerariesRequest struct {
    ItineraryID string `json:"itinerary_id"`
    Title       string `json:"title"`
    Description string `json:"description"`
}

type UpdateItinerariesResponse struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    Description string `json:"description"`
    StartDate  string `json:"start_date"`
    EndDate    string `json:"end_date"`
    AuthorID   string `json:"author_id"`
    UpdatedAt  string `json:"updated_at"`
}

type DeleteItinerariesRequest struct {
    ItineraryID string `json:"itinerary_id"`
}

type DeleteItinerariesResponse struct {
    MessageItinerary bool `json:"message_itinerary"`
}

type Itinerary struct {
    ID            string `json:"id"`
    Title         string `json:"title"`
    Author        Author `json:"author"`
    StartDate     string `json:"start_date"`
    EndDate       string `json:"end_date"`
    LikesCount    int32  `json:"likes_count"`
    CommentsCount int32  `json:"comments_count"`
    CreatedAt     string `json:"created_at"`
}

type GetAllItinerariesRequest struct {
    ItineraryID string `json:"itinerary_id"`
    Limit       int32  `json:"limit"`
    Offset      int32  `json:"offset"`
}

type GetAllItinerariesResponse struct {
    Itineraries []Itinerary `json:"itineraries"`
}

type ItinerariesFullInfoRequest struct {
    ItineraryID string `json:"itinerary_id"`
}

type ItinerariesFullInfoResponse struct {
    ID                string        `json:"id"`
    Title             string        `json:"title"`
    Description       string        `json:"description"`
    StartDate         string        `json:"start_date"`
    EndDate           string        `json:"end_date"`
    Author            Author        `json:"author"`
    Destinations      Destination   `json:"destinations"`
    LikesCount        int32         `json:"likes_count"`
    CommentsCount     int32         `json:"comments_count"`
    CreatedAt         string        `json:"created_at"`
    UpdatedAt         string        `json:"updated_at"`
}

type CommentItinerariesRequest struct {
    ItineraryID string `json:"itinerary_id"`
    Content     string `json:"content"`
    AuthorID    string `json:"author_id"`
}

type CommentItinerariesResponse struct {
    ID           string `json:"id"`
    Content      string `json:"content"`
    AuthorID     string `json:"author_id"`
    ItineraryID  string `json:"itinerary_id"`
    CreatedAt    string `json:"created_at"`
}

type GetDestinationsRequest struct {
    Country string `json:"country"`
    City    string `json:"city"`
    Limit   int32  `json:"limit"`
    Offset  int32  `json:"offset"`
}

type GetDestinationsResponse struct {
    ID                string   `json:"id"`
    Name              string   `json:"name"`
    Country           string   `json:"country"`
    Description       string   `json:"description"`
    PopularActivities []string `json:"popular_activities"`
}

type GetDestinationInfoRequest struct {
    DestinationID string `json:"destination_id"`
}

type GetDestinationInfoResponse struct {
    ID                   string   `json:"id"`
    Name                 string   `json:"name"`
    Country              string   `json:"country"`
    Description          string   `json:"description"`
    PopularActivities    []string `json:"popular_activities"`
    BestTimeToVisit      string   `json:"best_time_to_visit"`
    AverageCostPerDay    int32    `json:"average_cost_per_day"`
    Currency             string   `json:"currency"`
    Language             string   `json:"language"`
    TopAttractions       []string `json:"top_attractions"`
}

type SentMessageRequest struct {
    RecipientID string `json:"recipient_id"`
    SenderID    string `json:"sender_id"`
    Content     string `json:"content"`
}

type SentMessageResponse struct {
    ID          string `json:"id"`
    SenderID    string `json:"sender_id"`
    RecipientID string `json:"recipient_id"`
    Content     string `json:"content"`
    CreatedAt   string `json:"created_at"`
}

type Sender struct {
    ID       string `json:"id"`
    Username string `json:"username"`
}

type Recipient struct {
    ID       string `json:"id"`
    Username string `json:"username"`
}

type Message struct {
    ID        string    `json:"id"`
    Sender    Sender    `json:"sender"`
    Recipient Recipient `json:"recipient"`
    Content   string    `json:"content"`
    CreatedAt string    `json:"created_at"`
}

type GetAllMessagesRequest struct {
    UserID string `json:"user_id"`
    Limit  int32  `json:"limit"`
    Offset int32  `json:"offset"`
}

type GetAllMessagesResponse struct {
    Messages []Message `json:"messages"`
}

type CreateTravelTipRequest struct {
    Title    string `json:"title"`
    Content  string `json:"content"`
    Category string `json:"category"`
    AuthorID string `json:"author_id"`
}

type CreateTravelTipResponse struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    Content    string `json:"content"`
    Category   string `json:"category"`
    AuthorID   string `json:"author_id"`
    CreatedAt  string `json:"created_at"`
}

type Tip struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    Category   string `json:"category"`
    Author     Author `json:"author"`
    CreatedAt  string `json:"created_at"`
}

type GetTravelTipsRequest struct {
    Limit  int32 `json:"limit"`
    Offset int32 `json:"offset"`
}

type GetTravelTipsResponse struct {
    Tips []Tip `json:"tips"`
}

type UserStatisticsRequest struct {
    UserID string `json:"user_id"`
}

type MostPopularStory struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    LikesCount int32  `json:"likes_count"`
}

type MostPopularItinerary struct {
    ID         string `json:"id"`
    Title      string `json:"title"`
    LikesCount int32  `json:"likes_count"`
}

type UserStatisticsResponse struct {
    UserID             string             `json:"user_id"`
    TotalStories       int32              `json:"total_stories"`
    TotalItineraries   int32              `json:"total_itineraries"`
    TotalCountriesVisited int32           `json:"total_countries_visited"`
    TotalLikesReceived int32              `json:"total_likes_received"`
    TotalCommentsReceived int32           `json:"total_comments_received"`
    MostPopularStory   MostPopularStory   `json:"most_popular_story"`
    MostPopularItinerary MostPopularItinerary `json:"most_popular_itinerary"`
}

type GetTrendingDestinationsRequest struct{}

type GetTrendingDestinationsResponse struct {
    Destinations []Destination `json:"destinations"`
    Total        int32         `json:"total"`
}


// RegisterRequest represents the registration request payload.
type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email string    `json:"email" binding:"required"`
	FullName string    `json:"full_name" binding:"required"`
}

// LoginRequest represents the login request payload.
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// UpdateProfileRequest represents the update profile request payload.
type UpdateProfileRequest struct {
	UserID   string `json:"user_id"`
	Email    string `json:"email" binding:"required"`
	FullName string `json:"full_name" binding:"required"`
	Bio      string `json:"bio"`
}

// RefreshRequest represents the refresh token request payload.
type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

// Tokens represents the tokens response payload.
type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

// ProfileResponse represents the profile response payload.
type ProfileResponse struct {
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	FullName string `json:"full_name"`
	Bio      string `json:"bio"`
}

type Success struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type Failed struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}

type Logout struct{
	Message string `json:"message"`
}