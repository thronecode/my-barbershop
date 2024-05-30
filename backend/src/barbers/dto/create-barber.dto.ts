import { IsString, IsNotEmpty, IsUrl, IsNumber } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class CreateBarberDto {
  @ApiProperty({ description: 'Name of the barber' })
  @IsString()
  @IsNotEmpty()
  readonly name: string;

  @ApiProperty({ description: 'Photo url of the barber' })
  @IsUrl()
  @IsNotEmpty()
  readonly photo?: string;

  @ApiProperty({ description: 'Commission rate of the barber' })
  @IsNumber()
  @IsNotEmpty()
  commission_rate?: number;
}
