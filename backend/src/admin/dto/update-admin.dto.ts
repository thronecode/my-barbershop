import { IsString, IsNotEmpty, IsStrongPassword } from 'class-validator';
import { ApiProperty } from '@nestjs/swagger';

export class UpdateAdminDto {
  @ApiProperty({ description: 'Password of the admin' })
  @IsString()
  @IsStrongPassword()
  @IsNotEmpty()
  readonly password: string;
}
