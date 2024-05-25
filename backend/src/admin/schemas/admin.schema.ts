import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';

@Schema()
export class Admin extends Document {
  @Prop({ required: true })
  username: string;

  @Prop({ required: true })
  password: string;
}

@Schema()
export class AdminResponse extends Document {
  @Prop({ required: true })
  username: string;
}

export const AdminResponseSchema = SchemaFactory.createForClass(AdminResponse);
export const AdminSchema = SchemaFactory.createForClass(Admin);
