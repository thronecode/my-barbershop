import {
  Controller,
  Get,
  Post,
  Put,
  Delete,
  Body,
  Param,
  HttpStatus,
  UseGuards,
} from '@nestjs/common';
import { ServicesService } from './services.service';
import { CreateServiceDto } from './dto/create-service.dto';
import { UpdateServiceDto } from './dto/update-service.dto';
import { ApiTags, ApiResponse, ApiOperation } from '@nestjs/swagger';
import { Service } from './schemas/service.schema';
import { JwtAuthGuard } from '../auth/jwt-auth.guard';

@Controller('services')
@ApiTags('services')
export class ServicesController {
  constructor(private readonly servicesService: ServicesService) {}

  @UseGuards(JwtAuthGuard)
  @Get()
  @ApiOperation({ summary: 'List all services' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the list of all services',
    type: [Service],
  })
  async findAll(): Promise<Service[]> {
    return this.servicesService.findAll(false);
  }

  @UseGuards(JwtAuthGuard)
  @Get(':id')
  @ApiOperation({ summary: 'Get a service by id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the service',
    type: Service,
  })
  async findOne(@Param('id') id: string): Promise<Service> {
    return this.servicesService.findOne(id, false);
  }

  @UseGuards(JwtAuthGuard)
  @Post()
  @ApiOperation({ summary: 'Create a new service' })
  @ApiResponse({
    status: HttpStatus.CREATED,
    description: 'Return the created service',
    type: Service,
  })
  async create(@Body() createServiceDto: CreateServiceDto): Promise<Service> {
    return this.servicesService.create(createServiceDto);
  }

  @UseGuards(JwtAuthGuard)
  @Put(':id')
  @ApiOperation({ summary: 'Update a service by id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Return the updated service',
    type: Service,
  })
  async update(
    @Param('id') id: string,
    @Body() updateServiceDto: UpdateServiceDto,
  ): Promise<Service> {
    return this.servicesService.update(id, updateServiceDto);
  }

  @UseGuards(JwtAuthGuard)
  @Delete(':id')
  @ApiOperation({ summary: 'Delete a service by id' })
  @ApiResponse({
    status: HttpStatus.OK,
    description: 'Service successfully deleted',
    type: Service,
  })
  async remove(@Param('id') id: string): Promise<Service> {
    return this.servicesService.remove(id);
  }
}
