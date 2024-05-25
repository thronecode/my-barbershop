import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';
import { ApiProperty } from '@nestjs/swagger';

@Schema()
export class Admin extends Document {
  @Prop()
  @ApiProperty({ description: 'Id of the admin' })
  _id: string;

  @Prop({ required: true })
  @ApiProperty({ description: 'Username of the admin' })
  username: string;

  @Prop({ required: true })
  password: string;
}

export const AdminSchema = SchemaFactory.createForClass(Admin);
