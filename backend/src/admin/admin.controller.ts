import {
  Controller,
  Post,
  Get,
  Body,
  UseGuards,
  Param,
  Delete,
  Put,
  HttpStatus,
} from '@nestjs/common';
import { AdminService } from './admin.service';
import { CreateAdminDto } from './dto/create-admin.dto';
import { Admin } from './schemas/admin.schema';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';
import { UpdateAdminDto } from './dto/update-admin.dto';
import { GetJwtToken } from '../decorators/get-jwt-token.decorator';
import {
  ApiOperation,
  ApiProperty,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';

@Controller('admin')
export class AdminController {
  constructor(private readonly adminService: AdminService) {}

  @Post()
  @ApiOperation({ summary: 'Create a new admin' })
  @ApiTags('admin')
  @ApiProperty({ type: CreateAdminDto, description: 'Admin details' })
  @ApiResponse({
    status: HttpStatus.CREATED,
    description: 'Return the created admin',
    type: Admin,
  })
  async create(@Body() createAdminDto: CreateAdminDto): Promise<Admin> {
    return this.adminService.create(createAdminDto);
  }

  @Get()
  @ApiOperation({ summary: 'Get all admins' })
  @ApiTags('admin')
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return all admins',
    type: [Admin],
  })
  @UseGuards(JwtAuthGuard)
  async findAll(): Promise<Admin[]> {
    return this.adminService.findAll();
  }

  @Get(':id')
  @ApiOperation({ summary: 'Get admin by id' })
  @ApiTags('admin')
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the admin',
    type: Admin,
  })
  @UseGuards(JwtAuthGuard)
  async findOne(@Param('id') id: string): Promise<Admin> {
    return this.adminService.findOne(id);
  }

  @Delete(':id')
  @ApiOperation({ summary: 'Delete admin by id' })
  @ApiTags('admin')
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the deleted admin',
    type: Admin,
  })
  @UseGuards(JwtAuthGuard)
  async remove(
    @Param('id') id: string,
    @GetJwtToken() token: string,
  ): Promise<Admin> {
    return this.adminService.remove(id, token);
  }

  @Put(':id')
  @ApiOperation({ summary: 'Update admin by id' })
  @ApiTags('admin')
  @ApiProperty({ type: String, description: 'Admin id' })
  @ApiProperty({ type: UpdateAdminDto, description: 'Admin details' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the updated admin',
    type: Admin,
  })
  @UseGuards(JwtAuthGuard)
  async update(
    @Param('id') id: string,
    @Body() updateAdminDto: UpdateAdminDto,
  ): Promise<Admin> {
    return this.adminService.update(id, updateAdminDto);
  }
}
