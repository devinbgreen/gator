package main

import (
	"context"
	"fmt"
	"time"

	"github.com/bootdotdev/gator/internal/database"
	"github.com/google/uuid"
)

func handlerCreateFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}

	url := cmd.Args[0]

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	ff, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't create follow: %w", err)
	}

	fmt.Println("follow created successfully:")
	fmt.Printf(" * User:    %v\n", ff.UserName)
	fmt.Printf(" * Feed:    %v\n", ff.FeedName)
	return nil
}

func handlerListFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 0 {
		return fmt.Errorf("usage: %v", cmd.Name)
	}

	ff, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("couldn't list feeds: %w", err)
	}

	for _, f := range ff {
		fmt.Printf(" * Feed:    %v\n", f.FeedName)
	}

	return nil
}

func handlerDeleteFeedFollow(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 1 {
		return fmt.Errorf("usage: %v <url>", cmd.Name)
	}
	url := cmd.Args[0]
	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("couldn't find feed: %w", err)
	}

	//fmt.Printf(" * User:    %v\n", user.ID)
	//fmt.Printf(" * Feed:    %v\n", feed.ID)
	err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't delete follow: %w", err)
	}

	fmt.Printf("deleted feed:    %v\n", feed.Url)
	return nil
}
