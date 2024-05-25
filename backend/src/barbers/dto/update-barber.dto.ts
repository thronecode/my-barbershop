import {
  IsBoolean,
  IsNotEmpty,
  IsNumber,
  IsString,
  IsUrl,
} from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class UpdateBarberDto {
  @ApiProperty({ description: 'Name of the barber' })
  @IsString()
  @IsNotEmpty()
  readonly name?: string;

  @ApiProperty({ description: 'Photo url of the barber' })
  @IsUrl()
  @IsNotEmpty()
  readonly photo?: string;

  @ApiProperty({ description: 'Is the barber working today?' })
  @IsBoolean()
  @IsNotEmpty()
  readonly is_working?: boolean;

  @ApiProperty({ description: 'Commission rate of the barber' })
  @IsNumber()
  @IsNotEmpty()
  commission_rate?: number;
}
