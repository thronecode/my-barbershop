import { ApiProperty } from '@nestjs/swagger';
import {
  IsBoolean,
  IsEnum,
  IsInt,
  IsNotEmpty,
  IsNumber,
  IsString,
} from 'class-validator';

export class UpdateServiceDto {
  @ApiProperty({ description: 'Name of the service' })
  @IsString()
  @IsNotEmpty()
  name?: string;

  @ApiProperty({ description: 'Description of the service' })
  @IsString()
  @IsNotEmpty()
  description?: string;

  @ApiProperty({ description: 'Duration of the service in minutes' })
  @IsInt()
  @IsNotEmpty()
  duration?: number;

  @ApiProperty({ type: [String], description: 'Kinds of the service' })
  @IsString({ each: true })
  @IsEnum(['haircut', 'beard', 'eyebrows', 'other'])
  @IsNotEmpty({ each: true })
  kinds?: string[];

  @ApiProperty({ description: 'Is the service a combo?' })
  @IsBoolean()
  @IsNotEmpty()
  is_combo?: boolean;

  @ApiProperty({ description: 'Price of the service' })
  @IsNumber()
  @IsNotEmpty()
  price?: number;

  @ApiProperty({ description: 'Commission rate of the service' })
  @IsNumber()
  @IsNotEmpty()
  commission_rate?: number;
}
