import { Prop, Schema, SchemaFactory } from '@nestjs/mongoose';
import { Document } from 'mongoose';
import { ApiProperty } from '@nestjs/swagger';

@Schema()
export class Barber extends Document {
  @Prop()
  @ApiProperty({ description: 'Id of the barber' })
  _id: string;

  @Prop({ required: true })
  @ApiProperty({ description: 'Name of the barber' })
  name: string;

  @Prop()
  @ApiProperty({ description: 'Photo url of the barber' })
  photo?: string;

  @Prop({ default: false })
  @ApiProperty({ description: 'Is the barber working today?' })
  is_working: boolean;

  @Prop()
  @ApiProperty({ description: 'Commission rate of the barber' })
  commission_rate?: number;
}

export const BarberSchema = SchemaFactory.createForClass(Barber);
