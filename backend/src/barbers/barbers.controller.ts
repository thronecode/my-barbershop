import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Body,
  Param,
  UseGuards,
  HttpStatus,
} from '@nestjs/common';
import { BarbersService } from './barbers.service';
import { CreateBarberDto } from './dto/create-barber.dto';
import { UpdateBarberDto } from './dto/update-barber.dto';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';
import {
  ApiOperation,
  ApiProperty,
  ApiResponse,
  ApiTags,
} from '@nestjs/swagger';
import { Barber } from './schemas/barber.schema';

@Controller('barbers')
export class BarbersController {
  constructor(private readonly barbersService: BarbersService) {}

  @UseGuards(JwtAuthGuard)
  @Post()
  @ApiOperation({ summary: 'Create a new barber' })
  @ApiTags('barbers')
  @ApiProperty({ type: CreateBarberDto, description: 'Barber details' })
  @ApiResponse({
    status: HttpStatus.CREATED,
    description: 'Return the created barber',
    type: Barber,
  })
  async create(@Body() createBarberDto: CreateBarberDto) {
    return this.barbersService.create(createBarberDto);
  }

  @UseGuards(JwtAuthGuard)
  @Get()
  @ApiOperation({ summary: 'Get all barbers' })
  @ApiTags('barbers')
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return all barbers',
    type: [Barber],
  })
  async findAll() {
    return this.barbersService.findAll();
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  @ApiOperation({ summary: 'Get a barber by id' })
  @ApiTags('barbers')
  @ApiProperty({ type: String, description: 'Barber id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return a barber',
    type: Barber,
  })
  async findOne(@Param('id') id: string) {
    return this.barbersService.findOne(id);
  }

  @UseGuards(JwtAuthGuard)
  @Put(':id')
  @ApiOperation({ summary: 'Update a barber by id' })
  @ApiTags('barbers')
  @ApiProperty({ type: String, description: 'Barber id' })
  @ApiProperty({ type: UpdateBarberDto, description: 'Barber details' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the updated barber',
    type: Barber,
  })
  async update(
    @Param('id') id: string,
    @Body() updateBarberDto: UpdateBarberDto,
  ) {
    return this.barbersService.update(id, updateBarberDto);
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':id')
  @ApiOperation({ summary: 'Delete a barber by id' })
  @ApiTags('barbers')
  @ApiProperty({ type: String, description: 'Barber id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the deleted barber',
    type: Barber,
  })
  async remove(@Param('id') id: string) {
    return this.barbersService.remove(id);
  }
}
