import { Model, Document, Schema, model } from 'mongoose'

export interface IComment extends Document {
  author: string
  comment: string
}

const commentSchema = new Schema<IComment>(
  {
    author: { type: String, required: true },
    comment: { type: String, required: true },
  },
  { timestamps: true }
)

const Comment = model<IComment>('Comment', commentSchema)

export class CommentsRepository {
  private commentModel: Model<IComment>
  constructor(commentModel: Model<IComment>) {
    this.commentModel = commentModel
  }

  async fetchCommentById(id: string) {
    try {
      const comment: IComment | null = await this.commentModel.findById(
        id,
        '-_id -createdAt -updatedAt -__v'
      )
      return comment
    } catch (error) {
      throw new Error(`Error fetching comment with id: ${id} | ${error}`)
    }
  }

  async fetchAll() {
    try {
      const comments: IComment[] | null = await this.commentModel.find(
        {},
        '-_id -createdAt -updatedAt -__v'
      )
      return comments
    } catch (error) {
      throw new Error(`Error fetching comments: ${error}`)
    }
  }

  async storeComment(comment: Partial<IComment>) {
    const comm = new this.commentModel(comment)
    await comm.save()
  }

  async updateDoc(id: string, newComment: Partial<IComment>) {
    try {
      await this.commentModel.findByIdAndUpdate({ _id: id }, newComment)
    } catch (error) {
      throw new Error(`Error replacing comment with id: ${id} | ${error}`)
    }
  }

  async deleteDoc(id: string) {
    try {
      await this.commentModel.deleteOne({ _id: id })
    } catch (error) {
      throw new Error(`Error deliting comment with id: ${id} | ${error}`)
    }
  }
}

export const commentsRepo = new CommentsRepository(Comment)
