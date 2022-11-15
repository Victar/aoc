package adventofcode.year2018;

import adventofcode.BaseTest;
import lombok.Data;
import org.junit.Test;

import java.util.ArrayList;
import java.util.List;
import java.util.Set;

public class Day3 extends BaseTest {

	@Test public void runSilver() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day3/input.txt");
		List<Coord> coords = new ArrayList<>();
		for (final String input : data) {
			coords.add(new Coord(input));
		}
		int totalMatch = 0;
		for (int x = 0; x< 1200; x++) {
			for (int y = 0; y< 1200; y++) {
				int match = 0;
				for (Coord coord : coords){
					if (coord.inArea(x,y)){
						match ++;
					}
				}
				if ( match>=2){
					totalMatch++;
				}
			}
		}
		System.out.println(totalMatch);

	}

	@Test public void runGold() throws Exception {
		final ArrayList<String> data = readStringFromFile("year2018/day3/input.txt");
		List<Coord> coords = new ArrayList<>();
		for (final String input : data) {
			coords.add(new Coord(input));
		}
		int totalMatch = 0;

		for (Coord coord : coords){
			int match = 0;
			for (int x = coord.x+1; x<= coord.x + coord.sizeX; x++) {
				for (int y = coord.y+1; y<= coord.y + coord.sizeY; y++) {
					for (Coord c : coords){
						if (!c.equals(coord)){
							if (c.inArea(x,y)){
								match++;
							}
						}
					}

				}
			}
			if (match==0){
				System.out.println(match + " " +  coord);
			}
		}
	}



	@Data
	class Coord {
		String id;
		int x;
		int y;
		int sizeX;
		int sizeY;

		public Coord(String input){
			String in  = input.replace(": ",",").replace("x",",").replace(" @ ",",");
			String[] parts = in.split(",");
			id = parts[0];
			x = Integer.valueOf( parts[1]);
			y = Integer.valueOf( parts[2]);
			sizeX = Integer.valueOf( parts[3]);
			sizeY = Integer.valueOf( parts[4]);
		}

		public boolean inArea(int xIn, int yIn){
			return ( x <xIn && xIn <= (x+ sizeX) )&& (y <yIn && yIn <= (y+ sizeY));
		}
	}
}
