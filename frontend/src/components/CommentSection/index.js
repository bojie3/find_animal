import React from 'react'
import { 
  NewCommentContainer,
  CommentSectionContainer,
  CommentContainer,
  UserDetailsContainer,
  ProfilePic,
  Username,
  Comment
 } from './CommentSectionElements'
import UserProfilePic from '../../images/UserProfilePic.png'

const CommentSection = () => {
  return (
    <>
      <NewCommentContainer>
        New Comment here!
      </NewCommentContainer>
      <CommentSectionContainer>
        <CommentContainer>
          <UserDetailsContainer>
            <ProfilePic src={UserProfilePic} />
            <Username>Pengwing</Username>
          </UserDetailsContainer>
          <Comment>Penguins have knees!</Comment>
        </CommentContainer>
      </CommentSectionContainer>
    </>
  )
}

export default CommentSection